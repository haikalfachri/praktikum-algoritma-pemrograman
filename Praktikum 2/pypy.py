jumlah_matkul = int(input())
jumlah_sks = 0
ips = 0
ip = 0
while jumlah_matkul > 0:
    nilai, sks = input().split()
    sks = int(sks)
    if (nilai == "A" or nilai == "B" or nilai == "C" or nilai == "D" or nilai == "E") and sks > 0:
        if nilai == "A":
            ip += 4 * sks
            jumlah_sks += sks
        elif nilai == "B":
            ip += 3 * sks
            jumlah_sks += sks
        elif nilai == "C":
            ip += 2 * sks
            jumlah_sks += sks
        elif nilai == "D":
            ip += 1 * sks
            jumlah_sks += sks
        elif nilai == "E":
            ip += 0 * sks
            jumlah_sks += sks
        jumlah_matkul -= 1
    else:
        continue
ips = ip / jumlah_sks
print(ips)

    
