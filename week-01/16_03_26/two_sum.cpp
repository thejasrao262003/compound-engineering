class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int, int> prefix_sum;
        int max_val = INT_MIN;
        vector<int> ans;
        for(int i=0;i<nums.size();i++){
            if(prefix_sum.find(nums[i])!=prefix_sum.end()){
                ans.push_back(prefix_sum[nums[i]]);
                ans.push_back(i);
                return ans;
            }

            prefix_sum[target - nums[i]] = i;
        }
        return ans;
    }
};