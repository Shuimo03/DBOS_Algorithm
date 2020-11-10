def indexedSequentialSearch (arr,n,k):
    elements = [0]*20
    indices = [0]*20
    j,ind,start,end = 0,0,0,0
    set_flag = 0

    for i in range(0,n,3):
        # Storing element
        elements[ind] = arr[i]

        # Storing the index
        indices[ind] = i
        ind += 1

        if k < elements[0]:
            print("Not foud")
            exit(0)
        else:
            
            for i in range(1,ind+1):
                if k <= elements[i]:
                    start = indices[i-1]
                    end = indices[i]
                    set_flag = 1
                    break
            if set_flag == 0:
                start = indices[i-1]
                end = n
            for i in range(start,end+1):
                if k == arr[i]:
                    j = 1
                    break
            
            if j == 1:
                print("Found at index",i)
            else:
                print("Not found")

                
if __name__ == "__main__":
 
    arr = [6, 7, 8, 9, 10]
    n = len(arr)
 
    # Element to search
    k = 8
 
    # Function call
    indexedSequentialSearch(arr, n, k)