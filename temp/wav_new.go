package main

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func main() {
	// 定义一个包含所有WAV文件路径的数组
	files := []string{
		//"/Users/peanut/Desktop/w1.wav",
		"/Users/peanut/Desktop/w2.wav",
		"/Users/peanut/Desktop/w3.wav",
		// 添加更多文件路径...
	}

	// 创建一个新的WAV文件用于保存合并后的音频
	outFile, err := os.Create("output.wav")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// 创建一个新的PCM缓冲区用于保存合并后的音频数据
	outBuf := &audio.IntBuffer{
		Data:   []int{},
		Format: &audio.Format{},
	}

	bitDepth := 16 // 假设位深度为16，可以根据实际情况调整

	// 遍历每个文件
	for _, file := range files {
		// 判断音频采样率问题

		// 打开WAV文件
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer f.Close()

		// 读取WAV文件的头部信息和音频数据
		decoder := wav.NewDecoder(f)

		// 检查是否有错误或不支持的格式
		if !decoder.IsValidFile() {
			fmt.Println("Invalid WAV file:", file)
			return
		}

		buf, err := decoder.FullPCMBuffer()
		if err != nil {
			fmt.Println("Error reading PCM data from file:", err)
			return
		}

		// 简单地将多通道降为单通道（仅保留第一个通道）
		if buf.Format.NumChannels > 1 {
			buf.Data = downmixToMono(buf.Data, buf.Format.NumChannels)
			buf.Format.NumChannels = 1
		}

		// 将音频数据添加到输出缓冲区
		outBuf.Data = append(outBuf.Data, buf.Data...)
		outBuf.Format = buf.Format
	}

	// 创建一个新的WAV编码器并写入合并后的音频数据到输出文件中
	encoder := wav.NewEncoder(outFile, outBuf.Format.SampleRate, bitDepth, outBuf.Format.NumChannels, 1)
	err = encoder.Write(outBuf)
	if err != nil {
		fmt.Println("Error writing merged PCM data to output file:", err)
		return
	}

	duration := float64(len(outBuf.Data)) / float64(outBuf.Format.SampleRate)
	fmt.Printf("The duration of the merged audio is %f seconds.\n", duration)

	err = encoder.Close()
	if err != nil {
		fmt.Println("Error closing encoder:", err)
		return
	}

	fmt.Println("Successfully merged the WAV files into output.wav")
}

// downmixToMono 将多通道音频数据降为单通道，仅保留第一个通道的数据。
func downmixToMono(data []int, numChannels int) []int {
	monoData := make([]int, len(data)/numChannels)

	for i := range monoData {
		monoData[i] = data[i*numChannels]
	}

	return monoData
}
