package main

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func main() {
	// 打开第一个WAV文件
	file1, err := os.Open("/Users/peanut/Desktop/w1.wav")
	if err != nil {
		fmt.Println("Error opening file1:", err)
		return
	}
	defer file1.Close()

	// 打开第二个WAV文件
	file2, err := os.Open("/Users/peanut/Desktop/w3.wav")
	if err != nil {
		fmt.Println("Error opening file2:", err)
		return
	}
	defer file2.Close()

	// 创建一个新的WAV文件用于保存合并后的音频
	outFile, err := os.Create("output.wav")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// 读取第一个WAV文件的头部信息和音频数据
	decoder1 := wav.NewDecoder(file1)

	// 检查是否有错误或不支持的格式
	if !decoder1.IsValidFile() {
		fmt.Println("Invalid WAV file:", "file1.wav")
		return
	}

	buf1, err := decoder1.FullPCMBuffer()
	if err != nil {
		fmt.Println("Error reading PCM data from file1:", err)
		return
	}

	// 读取第二个WAV文件的头部信息和音频数据
	decoder2 := wav.NewDecoder(file2)

	if !decoder2.IsValidFile() {
		fmt.Println("Invalid WAV file:", "file2.wav")
		return
	}

	buf2, err := decoder2.FullPCMBuffer()
	if err != nil {
		fmt.Println("Error reading PCM data from file2:", err)
		return
	}

	// 检查采样率和通道数是否匹配，如果不匹配则进行简单的转换（例如：只取一部分通道）
	if buf1.Format.SampleRate != buf2.Format.SampleRate || buf1.Format.NumChannels != buf2.Format.NumChannels {
		fmt.Println("Sample rates or number of channels do not match between the two files.")

		// 简单地将多通道降为单通道（仅保留第一个通道）
		if buf1.Format.NumChannels > 1 {
			buf1.Data = downmixToMono_bak(buf1.Data, buf1.Format.NumChannels)
			buf1.Format.NumChannels = 1
		}
		if buf2.Format.NumChannels > 1 {
			buf2.Data = downmixToMono_bak(buf2.Data, buf2.Format.NumChannels)
			buf2.Format.NumChannels = 1
		}

		// 如果采样率不同，这里仅作提示，实际中可能需要重采样处理（可以使用第三方库如 SoX 或 PortAudio）
		if buf1.Format.SampleRate != buf2.Format.SampleRate {
			fmt.Println("Sample rates are different and need resampling. Please use a dedicated library for resampling.")
			return
		}
	}

	// 创建一个新的PCM缓冲区用于保存合并后的音频数据
	outBuf := &audio.IntBuffer{
		Data: make([]int, len(buf1.Data)+len(buf2.Data)),
		Format: &audio.Format{
			SampleRate:  buf1.Format.SampleRate,
			NumChannels: buf1.Format.NumChannels,
		},
	}

	copy(outBuf.Data, buf1.Data)
	copy(outBuf.Data[len(buf1.Data):], buf2.Data)

	bitDepth := 16 // 假设位深度为16，可以根据实际情况调整

	// 创建一个新的WAV编码器并写入合并后的音频数据到输出文件中
	encoder := wav.NewEncoder(outFile, outBuf.Format.SampleRate, bitDepth, outBuf.Format.NumChannels, 1)
	err = encoder.Write(outBuf)
	if err != nil {
		fmt.Println("Error writing merged PCM data to output file:", err)
		return
	}

	err = encoder.Close()
	if err != nil {
		fmt.Println("Error closing encoder:", err)
		return
	}

	fmt.Println("Successfully merged the WAV files into output.wav")
}

// downmixToMono 将多通道音频数据降为单通道，仅保留第一个通道的数据。
func downmixToMono_bak(data []int, numChannels int) []int {
	monoData := make([]int, len(data)/numChannels)

	for i := range monoData {
		monoData[i] = data[i*numChannels]
	}

	return monoData
}
