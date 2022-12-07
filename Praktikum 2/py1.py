status = False
i = 1
ronde = 1
pemenang = "A"
kalah = "B"
angka_ajaib = -101
while not status:
    print("Ronde ke-{}: ".format(ronde))
    angka_rahasia = int(input("{} - masukkan angka rahasia: ".format(pemenang)))
    if angka_rahasia == angka_ajaib:
        status = True
        print("Permainan selesai")
    else:
        ronde += 1
        while i <= 3:
            tebakan = int(input("{} - masukkan angka tebakan ke-{}: ".format(kalah, i)))
            if tebakan == angka_rahasia:
                if pemenang == "A":
                    pemenang = "B"
                    kalah = "A"
                elif pemenang == "B":
                    pemenang = "A"
                    kalah = "B"
                print(pemenang, "adalah pemenangnya")
                break
            else:
                i += 1
                if i > 3:
                    if pemenang == "B":
                        pemenang = "B"
                        kalah = "A"
                    elif pemenang == "A":
                        pemenang = "A"
                        kalah = "B"
                    print(pemenang, "adalah pemenangnya")
        i = 1
                    

    