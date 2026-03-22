// Passed 40/40 test cases.
class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        int n = nums.size();
        k = k % n;

        reverse(nums.begin(), nums.end());
        reverse(nums.begin(), nums.begin() + k);
        reverse(nums.begin() + k, nums.end());
    }
};

// Passed 39/40 test cases. Time limit exceeded for the last test case.
class Solution {
public:
    void rotate(vector<int>& nums, int k) {
        int n = nums.size();
        k = k%n;
        for(int i = 0;i<k;i++){
            int temp = nums[n-1];
            for(int i=n-2;i>=0;i--){
                nums[i+1] = nums[i];
            }
            nums[0] = temp;
        }
    }
};