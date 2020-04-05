#include<stdio.h>
#include<stdlib.h>


//primo incluso secondo escluso
int fmax_my(int a, int b){
  if(a > b)
    return a;
  else
    return b;
}

int elem(int *arr, int inizio, int fine, int n1){
  int return_value = 0,half = 0, n = n1, out1 = 0, out2 = 0, max_out = 0;
  /*printf("input\n");
  for(int i = inizio; i < fine-1; i++ ){
    printf("%d ",arr[i]);
  }
  printf("\n");*/
  if(n1 == 1){
    //printf("ritorno\n");
    return 1;
  }
  for(int i = inizio; i < fine; i++){
    if(arr[i] > arr[i+1]){
      half = (inizio + fine) / 2;
      printf("chiamata1 %d %d %d \n", inizio, half-1, half);
      out1 = elem(arr, inizio, half, half);
      printf("chiamata2 %d %d %d \n", half, fine, half);
      out2 = elem(arr, half, fine, half);
      max_out = fmax_my(out1, out2);
      return max_out;
    } else {
      return_value = n1;
    }
  }
  return return_value;
}

int main(){
  int n, return_value1 = 0, in;
  int *arr;
  scanf("%d", &n);
  arr = calloc(n, sizeof(int));
  for(int i = 0; i < n; i++){
    scanf(" %d", &in);
    arr[i] = in;
  }
  return_value1 = elem(arr, 0, n, n);
  printf("%d\n", return_value1);
}

/*
16
78 16 40 90 4 13 99 19 41 47 50 68 71 79 96 76
*/
