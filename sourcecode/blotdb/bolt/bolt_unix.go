// +build !windows,!plan9,!solaris

package bolt

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

// flock acquires an advisory lock on a file descriptor.
func flock(db *DB, mode os.FileMode, exclusive bool, timeout time.Duration) error {
	var t time.Time
	for {
		// If we're beyond our timeout then return an error.
		// This can only occur after we've attempted a flock once.
		if t.IsZero() {
			t = time.Now()
		} else if timeout > 0 && time.Since(t) > timeout {
			return ErrTimeout
		}
		flag := syscall.LOCK_SH
		if exclusive {
			flag = syscall.LOCK_EX
		}

		// Otherwise attempt to obtain an exclusive lock.
		// http://c.biancheng.net/view/5730.html#:~:text=syscall.Flock%20%28int%20%28f.Fd%20%28%29%29%2C%20syscall.LOCK_EX%7Csyscall.LOCK_NB%29%20%E6%88%91%E4%BB%AC%E9%87%87%E7%94%A8%E4%BA%86%20syscall.LOCK_EX%E3%80%81syscall.LOCK_NB%EF%BC%8C%E8%BF%99%E6%98%AF%E4%BB%80%E4%B9%88%E6%84%8F%E6%80%9D%E5%91%A2%EF%BC%9F.%20flock,flock%20%E5%B0%86%E6%96%87%E4%BB%B6%E9%94%81%E4%BD%8F%EF%BC%8C%E5%8F%A6%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E5%8F%AF%E4%BB%A5%E7%9B%B4%E6%8E%A5%E6%93%8D%E4%BD%9C%E6%AD%A3%E5%9C%A8%E8%A2%AB%E9%94%81%E7%9A%84%E6%96%87%E4%BB%B6%EF%BC%8C%E4%BF%AE%E6%94%B9%E6%96%87%E4%BB%B6%E4%B8%AD%E7%9A%84%E6%95%B0%E6%8D%AE%EF%BC%8C%E5%8E%9F%E5%9B%A0%E5%9C%A8%E4%BA%8E%20flock%20%E5%8F%AA%E6%98%AF%E7%94%A8%E4%BA%8E%E6%A3%80%E6%B5%8B%E6%96%87%E4%BB%B6%E6%98%AF%E5%90%A6%E8%A2%AB%E5%8A%A0%E9%94%81%EF%BC%8C%E9%92%88%E5%AF%B9%E6%96%87%E4%BB%B6%E5%B7%B2%E7%BB%8F%E8%A2%AB%E5%8A%A0%E9%94%81%EF%BC%8C%E5%8F%A6%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E5%86%99%E5%85%A5%E6%95%B0%E6%8D%AE%E7%9A%84%E6%83%85%E5%86%B5%EF%BC%8C%E5%86%85%E6%A0%B8%E4%B8%8D%E4%BC%9A%E9%98%BB%E6%AD%A2%E8%BF%99%E4%B8%AA%E8%BF%9B%E7%A8%8B%E7%9A%84%E5%86%99%E5%85%A5%E6%93%8D%E4%BD%9C%EF%BC%8C%E4%B9%9F%E5%B0%B1%E6%98%AF%E5%BB%BA%E8%AE%AE%E6%80%A7%E9%94%81%E7%9A%84%E5%86%85%E6%A0%B8%E5%A4%84%E7%90%86%E7%AD%96%E7%95%A5%E3%80%82.%20flock%20%E4%B8%BB%E8%A6%81%E4%B8%89%E7%A7%8D%E6%93%8D%E4%BD%9C%E7%B1%BB%E5%9E%8B%EF%BC%9A.%20LOCK_SH%EF%BC%9A%E5%85%B1%E4%BA%AB%E9%94%81%EF%BC%8C%E5%A4%9A%E4%B8%AA%E8%BF%9B%E7%A8%8B%E5%8F%AF%E4%BB%A5%E4%BD%BF%E7%94%A8%E5%90%8C%E4%B8%80%E6%8A%8A%E9%94%81%EF%BC%8C%E5%B8%B8%E8%A2%AB%E7%94%A8%E4%BD%9C%E8%AF%BB%E5%85%B1%E4%BA%AB%E9%94%81%EF%BC%9B.%20LOCK_EX%EF%BC%9A%E6%8E%92%E4%BB%96%E9%94%81%EF%BC%8C%E5%90%8C%E6%97%B6%E5%8F%AA%E5%85%81%E8%AE%B8%E4%B8%80%E4%B8%AA%E8%BF%9B%E7%A8%8B%E4%BD%BF%E7%94%A8%EF%BC%8C%E5%B8%B8%E8%A2%AB%E7%94%A8%E4%BD%9C%E5%86%99%E9%94%81%EF%BC%9B.
		//flock 主要三种操作类型：
		//LOCK_SH：共享锁，多个进程可以使用同一把锁，常被用作读共享锁；
		//LOCK_EX：排他锁，同时只允许一个进程使用，常被用作写锁；
		//LOCK_UN：释放锁。
		//
		//进程使用 flock 尝试锁文件时，如果文件已经被其他进程锁住，进程会被阻塞直到锁被释放掉，或者在调用 flock 的时候，
		//采用 LOCK_NB 参数。在尝试锁住该文件的时候，发现已经被其他服务锁住，会返回错误，错误码为 EWOULDBLOCK。
		err := syscall.Flock(int(db.file.Fd()), flag|syscall.LOCK_NB)
		if err == nil {
			return nil
		} else if err != syscall.EWOULDBLOCK {
			return err
		}

		// Wait for a bit and try again.
		time.Sleep(50 * time.Millisecond)
	}
}

// funlock releases an advisory lock on a file descriptor.
func funlock(db *DB) error {
	return syscall.Flock(int(db.file.Fd()), syscall.LOCK_UN)
}

// mmap memory maps a DB's data file.
func mmap(db *DB, sz int) error {
	// Map the data file to memory.
	b, err := syscall.Mmap(int(db.file.Fd()), 0, sz, syscall.PROT_READ, syscall.MAP_SHARED|db.MmapFlags)
	if err != nil {
		return err
	}

	// Advise the kernel that the mmap is accessed randomly.
	if err := madvise(b, syscall.MADV_RANDOM); err != nil {
		return fmt.Errorf("madvise: %s", err)
	}

	// Save the original byte slice and convert to a byte array pointer.
	db.dataref = b
	db.data = (*[maxMapSize]byte)(unsafe.Pointer(&b[0]))
	db.datasz = sz
	return nil
}

// munmap unmaps a DB's data file from memory.
func munmap(db *DB) error {
	// Ignore the unmap if we have no mapped data.
	if db.dataref == nil {
		return nil
	}

	// Unmap using the original byte slice.
	err := syscall.Munmap(db.dataref)
	db.dataref = nil
	db.data = nil
	db.datasz = 0
	return err
}

// NOTE: This function is copied from stdlib because it is not available on darwin.
func madvise(b []byte, advice int) (err error) {
	_, _, e1 := syscall.Syscall(syscall.SYS_MADVISE, uintptr(unsafe.Pointer(&b[0])), uintptr(len(b)), uintptr(advice))
	if e1 != 0 {
		err = e1
	}
	return
}
