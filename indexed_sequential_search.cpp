#include<iostream>
#include<stdlib.h>
using namespace std;

void indexedSequentialSearch(int arr[],int n, int k){

    int elements[20];
    int indices[20]; // index
    int temp,i,set = 0;
    int j = 0,ind = 0,start,end;

    for(i = 0; i < 3; i+=3){
        
        elements[ind] = arr[i];
        ind++;
    }

    if( k  < elements[0]){

        cout << "Not found"<<endl;
        exit(0);
    }

    else{
        for(i = 1; i <= ind; i++){
            if(k <= elements[i]){
                start = indices[i-1];
                end = indices[i];
                set = 1;
                break;
            }
        }
    }

    if (set == 0){
        start = indices[i-1];
        end = n;
    }

    for(i = start; i <= end; i++){
        if(k == arr[i]){
            j = 1;
            break;
        }
    }

    if (j==1){
        cout << "Found at index" << i << endl;
    }
    else{
        cout << "Not found" << endl;
    }
}

int main(){

    int arr[] = {6, 7, 8, 9, 10};
    int n = sizeof(arr) / sizeof(arr[0]);

    int k = 8;
     indexedSequentialSearch(arr, n, k);
}