import csv

def rank(url_set, url_link):
    for url in url_set:
        link_to_me = list(filter(lambda x: x[1] == url, url_link))
        for link in link_to_me:
            url_set[url] += url_set[link[0]]*0.85
            url_set[link[0]] *= 0.15

def iterate_rank(n, set, link):
    for _ in range(0,n):
        rank(set, link)

def read_csv_return_tuple_array(path):
    result = []
    with open(path, newline='') as csvfile:
        reader = csv.reader(csvfile, delimiter=';')
        next(reader)
        for row in reader:
            result.append([row[0], row[1]])
    return result