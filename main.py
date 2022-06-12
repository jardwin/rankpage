import csv

def rank(url_set, url_link):
    for url in url_set:
        link_to_me = list(filter(lambda x: x[1] == url, url_link))
        for link in link_to_me:
            url_set[url] += (url_set[link[0]]*0.85)/len(list(filter(lambda x: x[0] == link[0], url_link)))

def iterate_rank(n, set, link):
    for _ in range(0,n):
        rank(set, link)
        reduce_rank(set, link)
        if len(link)!=0:
            distribut_rank(set, link)

def reduce_rank(set, links):
    filtered = []
    for link in links:
        if next((x for x in filtered if x == link[0]), None) is None:
            filtered.append(link[0])

    for page in filtered:
        set[page]=set[page]*0.15

def distribut_rank(set, links):
    def toto(x):
        val = set[x[0]]
        set[x[0]]=0
        return val
    rest = sum(list(map(toto, links)))
    for page in set:
        set[page]+=rest/len(set)

def create_set_from_link(links):
    filtered = {}
    for link in links:
        if next((x for x in filtered if x == link[0]), None) is None:
            filtered[link[0]] = 0
        if next((x for x in filtered if x == link[1]), None) is None:
            filtered[link[0]] = 0

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