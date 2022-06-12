import csv

def rank(url_set):
    for url in url_set:
        url_set[url]+=url_set[url]*0.15

mySet = {
    "toto.com":0.33,
    "titi.com":0.33,
    "tutu.com":0.33
}

for x in range(0,5):
    rank(mySet)

for url in mySet:
    print(url, "to", mySet[url])
