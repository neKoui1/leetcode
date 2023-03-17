#include<iostream>
#include<iomanip>
#include<cmath>
// #include<vector>

using namespace std;

int main(){
    int exce = 0, pass = 0;
    double exceRatio = 0, passRatio = 0;
    // vector<int> vec;
    int num = 0;
    cin >> num;
    for ( int i = 0; i < num; i ++ ){
        int input = 0;
        cin >> input;
        
        if ( input >= 60 ){
            pass ++;
        }
        if ( input >= 85 ) {
            exce ++;
        }
    }
    exceRatio = exce / double(num);
    passRatio = pass / double(num);

    // if ( exceRatio - int(exceRatio) >= 0.5 ) {
    //     exceRatio ++;
    // }
    // if ( passRatio - int(exceRatio) >= 0.5 ) {
    //     passRatio ++;
    // }

    cout << round(passRatio * 100) << '%' << endl;
    cout << round(exceRatio * 100) << '%' << endl;


    return 0;
}