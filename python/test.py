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
        self.assertEqual(0.0375, data['toto.com'])
        self.assertEqual(0.9625, data['titi.com'])

    def test_rank_two_entry_with_ref_three_iteration(self):
        data = {"toto.com":0.5,"titi.com":0.5}
        main.iterate_rank(3, data, [["toto.com", "titi.com"]])
        self.assertEqual(0.0002109375, data['toto.com'])
        self.assertEqual(0.9997890625, data['titi.com'])

    def test_rank_three_entry_with_ref(self):
        data = {"toto.com":1/3,"titi.com":1/3,"tata.com":1/3}
        main.iterate_rank(1, data, [["toto.com", "titi.com"],["toto.com", "tata.com"]])
        self.assertEqual(0.016666666666666666, data['toto.com'])
        self.assertEqual(0.49166666666666664, data['tata.com'])
        self.assertEqual(0.49166666666666664, data['titi.com'])
        
    def test_rank_five_entry_with_ref(self):
        data = {"toto.com":0.2,"titi.com":0.2,"tutu.com":0.2,"tata.com":0.2,"tete.com":0.2}
        main.iterate_rank(1, data, [["titi.com", "toto.com"],["tutu.com", "toto.com"],["tata.com", "toto.com"]])
        self.assertEqual(0.7280000000000001, data['toto.com'])
        self.assertEqual(0.018, data['titi.com'])
        self.assertEqual(0.018, data['tutu.com'])
        self.assertEqual(0.018, data['tata.com'])
        self.assertEqual(0.218, data['tete.com'])

    def test_csv_reading(self):
        result = main.read_csv_return_tuple_array("../testdataset.csv")
        self.assertEqual("http://toto.com", result[0][0])
        self.assertEqual("http://titi.com", result[0][1])

    def test_create_set_from_link(self):
        set = main.create_set_from_link([["AAA", "BBB"],["BBB", "CCC"],["AAA", "CCC"]])
        self.assertEqual(1/3,set["AAA"])
        self.assertEqual(1/3,set["BBB"])
        self.assertEqual(1/3,set["CCC"])
        
    def test_full(self):
        result = main.read_csv_return_tuple_array("../dataset.csv")
        data = main.create_set_from_link(result)
        main.iterate_rank(1, data, result)
        self.assertEqual(0.00023853366999338602,data["CGO"])

if __name__ == '__main__':
    unittest.main()