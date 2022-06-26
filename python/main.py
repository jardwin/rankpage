import csv
import sys

def rank(url_set, url_link):
    for url in url_set:
        for link in findLinkToMe(url_link, url):
            url_set[url] += (url_set[link[0]]*0.85)/countLinkFromMe(url_link, link[0])

def findLinkToMe(urlLinks, me):
    linked = []
    for url in urlLinks:
        if url[1] == me and me not in linked:
            linked.append(url)
    return linked

def countLinkFromMe(urlLinks, me):
    cmp = 0
    for url in urlLinks:
        if url[0] == me:
            cmp+=1
    return cmp


def iterate_rank(n, set, link):
    for _ in range(0,n):
        rank(set, link)
        reduce_rank(set, link)
        if len(link)!=0:
            distribut_rank(set, link)

def reduce_rank(set, links):
    alreadyDone = []
    for link in links:
        if link[0] not in alreadyDone:
            set[link[0]]=set[link[0]]*0.15
            alreadyDone.append(link[0])

def distribut_rank(set, links):
    rest = 0
    for link in links:
        rest+=set[link[0]]
        set[link[0]]=0
    for page in set:
        set[page]+=rest/len(set)

def create_set_from_link(links):
    filtered = {}
    for link in links:
        if link[0] not in filtered:
            filtered[link[0]] = 0
        if link[1] not in filtered:
            filtered[link[1]] = 0

    for page in filtered:
        filtered[page]=1/len(filtered)
    return filtered

def read_csv_return_tuple_array(path):
    result = []
    with open(path, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=';')
        next(reader)
        for row in reader:
            result.append([row[0], row[1]])
    return result

if len(sys.argv) > 2:
    result = read_csv_return_tuple_array(sys.argv[1])
    data = create_set_from_link(result)
    iterate_rank(int(sys.argv[2]), data, result)
    # for link, score in data.items():
    #     print("{:<8} {:<15}".format(link, score))
    with open("./result.csv", "w") as f:
        writer = csv.writer(f, delimiter=';')
        writer.writerow(["url", "note"])
        for page in data:
            writer.writerow([page, data[page]])
