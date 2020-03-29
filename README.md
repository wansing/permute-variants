# permute-variants

Create all permutations of words, where each word can have different variants (spellings, notations).

## Build

```
git clone https://github.com/wansing/permute-variants
go build
```

## Usage

```
./permute-variants file1.txt file2.txt file3.txt ...
```

## Example

```
$ cat a.txt
apple
banana
cherry

$ cat b.txt
apartment
house
tent

$ ./permute-variants a.txt b.txt
appleapartment
applehouse
appletent
bananaapartment
bananahouse
bananatent
cherryapartment
cherryhouse
cherrytent
apartmentapple
houseapple
tentapple
apartmentbanana
housebanana
tentbanana
apartmentcherry
housecherry
tentcherry
```
