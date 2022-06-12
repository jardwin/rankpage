import unittest

import main


class TestUtils(unittest.TestCase):
    def test_rank_one_entry(self):
        data = {"toto.com":1}
        main.iterate_rank(1, data, [])
        self.assertEqual(1, data['toto.com'])

    def test_rank_two_entry_unref(self):
        data = {"toto.com":0.5,"titi.com":0.5}
        main.iterate_rank(1, data, [])
        self.assertEqual(0.5, data['toto.com'])
        self.assertEqual(0.5, data['titi.com'])

    def test_rank_two_entry_with_ref(self):
        data = {"toto.com":0.5,"titi.com":0.5}
        main.iterate_rank(1, data, [["toto.com", "titi.com"]])
        self.assertEqual(0.075, data['toto.com'])
        self.assertEqual(0.925, data['titi.com'])

    def test_csv_reading(self):
        result = main.read_csv_return_tuple_array("./testdataset.csv")
        self.assertEqual("http://toto.com", result[0][0])
        self.assertEqual("http://titi.com", result[0][1])

if __name__ == '__main__':
    unittest.main()