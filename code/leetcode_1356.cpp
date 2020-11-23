/**
 * @date    2020-11-06 ‏‎16:15:32
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @link 	github.com/taseikyo
 */

class Solution {
  public:
	vector<int> sortByBits(vector<int>& arr) {
		vector<int> ret;
		map<int, vector<int>> mp;
		for (auto& num : arr) {
			int cnt = 0;
			int tmp = num;
			while (tmp != 0) {
				if (tmp & 1)
					cnt++;
				tmp >>= 1;
			}
			mp[cnt].push_back(num);
		}
		for (auto& iter : mp) {
			sort(iter.second.begin(), iter.second.end());
			for (auto& it : iter.second)
				ret.push_back(it);
		}
		return ret;
	}
};
