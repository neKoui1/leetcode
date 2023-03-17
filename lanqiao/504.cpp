// 小蓝正在学习一门神奇的语言，这门语言中的单词都是由小写英文字母组成，
// 有些单词很长，远远超过正常英文单词的长度。小蓝学了很长时间也记不住一些单词，
// 他准备不再完全记忆这些单词，而是根据单词中哪个字母出现得最多来分辨单词。
// 现在，请你帮助小蓝，给了一个单词后，帮助他找到出现最多的字母和这个字母出现的次数。

/**
 * 第一眼用哈希
 * 可惜没怎么用过cpp的哈希表
 * CSDN
*/
#include<iostream>
#include<map>
#include<string>

using namespace std;

int main(){
    map<char, int> mapS;
    int max = 0;
    string input = "";
    char res = ' ';
    cin >> input;
    for ( int i = 0 ; i < input.size(); i++ ){
        if ( mapS.count(input[i]) == 0 ){
            mapS[input[i]] = 1;
        }else {
            mapS[input[i]] ++;
        }
    }

    map<char, int> :: iterator it = mapS.begin();
    map<char, int> :: iterator itEnd = mapS.end();
    while(it!=itEnd){
        if( it -> second > max) {
            max = it -> second;
            res = it -> first;

        } else if ( it -> second == max && it -> first < res ){
            res = it -> first;
        }
        it ++;
    }
    cout << res << endl;
    cout << max << endl;

    return 0;
}
