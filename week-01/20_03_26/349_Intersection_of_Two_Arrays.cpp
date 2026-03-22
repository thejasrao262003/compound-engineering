class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        unordered_set<int> s1(nums1.begin(), nums1.end());
        unordered_set<int> result;

        for(auto i:nums2){
            if(s1.count(i)!=0){
                result.insert(i);
                s1.erase(i);
            }
            
        }
        return vector<int>(result.begin(), result.end());
    }
};