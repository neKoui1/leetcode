#include<iostream>
using namespace std;

int main(){
    int a=4;
    int *p=&a;
    int *q=p;
    cout<<*p<<endl;
    cout<<*q<<endl;
    *p=10;
    cout<<a<<endl;
    cout<<*p<<endl;
    cout<<*q<<endl;

    return 0;
}