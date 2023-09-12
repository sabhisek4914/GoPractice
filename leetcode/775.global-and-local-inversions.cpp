/*
 * @lc app=leetcode id=775 lang=cpp
 *
 * [775] Global and Local Inversions
 */

// @lc code=start
class Solution {
public:
    vector<int> temp;// To Avoid TLE
    long long merge(vector<int>& arr,int s,int e){
        temp.clear(); //Clearing it every time due to global temp
        int i=s;
        int m=(s+e)/2;
        int j= m+1;
        long long cnt=0;
        
        while(i<=m && j<=e){
            if(arr[i]<arr[j]){
                temp.push_back(arr[i]);
                i++;
            }else{
                cnt+=m-i+1;
                temp.push_back(arr[j]);//sa
                j++;
            }
        }
        while(i<=m){
            temp.push_back(arr[i++]);
        }
        while(j<=e){
            temp.push_back(arr[j++]);
        }
        int k=0;
        for(int i=s;i<=e;i++){
            arr[i]=temp[k++];
        }
        // cout<<s<<" "<<e<<" "<<cnt<<endl;
        return cnt;
    }
    long long mergeSort(vector<int>& arr,int s, int e){
        if(s>=e)
            return 0;
        int m=(s+e)/2;
        long long c1=0;
        // int c1=mergeSort(arr,s,m);
        // int c2=mergeSort(arr,m+1,e);
        // int c3=merge(arr,s,e);
        // return c1+c2+c3;
        c1+=mergeSort(arr,s,m);
        c1+=mergeSort(arr,m+1,e);
        c1+=merge(arr,s,e);
        return c1;
    }
    bool isIdealPermutation(vector<int>& nums) {
        bool check=false;
        
        long long localCount=0;
        long long globalCount=0;

        for(int i=0;i<nums.size()-1;i++){
            if(nums[i]>nums[i+1])
                localCount++;
        }
        globalCount=mergeSort(nums,0,nums.size()-1);

        cout<<localCount<<" "<<globalCount<<endl;
        if(globalCount==localCount) return true;
        return check;
    }
};
// @lc code=end

