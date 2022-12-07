package main

import "fmt"
import "sort"

func main(){
	const NMAX = 100
	type Provinsi struct{
		nama string
		populasi int
		tumbuh float64
	}
	type DataProvinsi struct{
		tabel [NMAX]*Provinsi
		nProvinsi int
	}
	
	var info Provinsi
	dataProv := make([]string, 0)
	dataPopulasi := make([]int, 0)
	dataRate := make([]float64, 0)
	
	fmt.Println("Data pertumbuhan provinsi:")
	fmt.Scan(&info.nama, &info.populasi, &info.tumbuh)
	for info.nama != "NONE" || info.populasi != 0 || info.tumbuh != 0.0{
		dataProv = append(dataProv, info.nama)
		dataPopulasi = append(dataPopulasi, info.populasi)
		dataRate = append(dataRate, info.tumbuh)
		fmt.Scan(&info.nama, &info.populasi, &info.tumbuh)
	}

	var cariProv string
	var posisi int
	fmt.Print("Nama provinsi? ")
	fmt.Scan(&cariProv)
	for i:=0;i<len(dataProv);i++{
		if dataProv[i] == cariProv{
			posisi = i
		}
	}
	fmt.Printf("%v %v %v\n", cariProv, dataPopulasi[posisi], dataRate[posisi])

	var tempPopulasi float64
	var populasiTahunDepan int
	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scan(&cariProv)
	for i:=0;i<len(dataProv);i++{
		if dataProv[i] == cariProv{
			posisi = i
		}
	}
	tempPopulasi = float64(dataPopulasi[posisi]) + (float64(dataPopulasi[posisi]) * dataRate[posisi])
	populasiTahunDepan = int(tempPopulasi)
	fmt.Printf("Populasi Provinsi %v tahun depan: %v\n", cariProv, populasiTahunDepan)
	
	temp:=make([]int, 0)
	for i:=0;i<len(dataPopulasi);i++{
		temp = append(temp, dataPopulasi[i])
	}
	dataPopulasiUrut := dataPopulasi
	sort.Ints(dataPopulasiUrut) 

	dataProvUrut := make([]string,0)
	dataRateUrut := make([]float64, 0)
	for i:=0;i<len(dataPopulasiUrut);i++{
		for j:=0;j<len(dataPopulasi);j++{ 
			if dataPopulasiUrut[i] == temp[j]{
				dataProvUrut = append(dataProvUrut, dataProv[j])
				dataRateUrut = append(dataRateUrut, dataRate[j])
				break
			}
		}
	}
	fmt.Println("Urutan data pertumbuhan provinsi berdasarkan populasi terurut menurun:")
	var idx int
	idx = len(dataPopulasiUrut)-1
	for idx > -1{
		fmt.Println(dataProvUrut[idx], dataPopulasiUrut[idx], dataRateUrut[idx])
		idx--
	}
}
