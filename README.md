# fabric-bank

Projekat iz predmeta PDASP - 2023/2024 - grupa G6.

Autori:
- Aleksa Bajat (E2 114/2023)
- Aleksandar Radišić (E2 17/2023)
- Bojan Popržen (E2 4/2022)

## Pokretanje

```bash
# Preuzmi potrebne Hyperledger fabric docker slike i preuzmi binarne fajlove i
# prebaci ih u ./fabric-samples/bin i ./fabric-samples/config.
./install-fabric.sh --fabric-version 2.2.6

# Spusti prethodno podignutu mrezu i podigni novu sa 4 organizacije sa po 4
# peer-a.
cd ./fabric-samples/test-network
./network.sh down
./network.sh up
````