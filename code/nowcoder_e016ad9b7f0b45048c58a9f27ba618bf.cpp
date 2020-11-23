/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2020-11-10 20:40:05
 * @link    github.com/taseikyo
 */

class Finder {
  public:
	int findKth(vector<int> a, int n, int K) {
		return findKth(a, 0, n - 1, n - K + 1);
	}
	int findKth(vector<int>&a, int l, int r, int K) {
		if (l == r) return a[l];
		int i = l - 1, j = r + 1;

		int x = a[l + r >> 1];
		while (i < j) {
			do i++; while (a[i] < x);
			do j--; while (x < a[j]);
			if (i < j) swap(a[i], a[j]);
		}
		int s = j - l + 1;
		if (K <= s) return findKth(a, l, j, K);
		else return findKth(a, j + 1, r, K - s);
	}
};
