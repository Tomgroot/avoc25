#include <stdio.h>
#include <string.h>
#include <time.h>

long long global_counter = 0;

int is_symmetric_split(long number) {
    char str[64];
    
    int len = snprintf(str, sizeof(str), "%ld", number);
    
    if (len % 2 != 0) {
        return 0;
    }
    
    int half_len = len / 2;
    
    if (strncmp(str, str + half_len, half_len) == 0) {
        return 1;
    }
    
    return 0;
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <input_file>\n", argv[0]);
        return 1;
    }

    FILE *file = fopen(argv[1], "r");
    if (!file) {
        perror("Error opening file");
        return 1;
    }

    long min, max;
    clock_t start_time = clock();
    
    while (fscanf(file, "%ld-%ld", &min, &max) == 2) {
        
        for (long i = min; i <= max; i++) {
            if (is_symmetric_split(i)) {
                global_counter += i;
            }
        }

        // Next character should be a comma or EOF
        char separator = fgetc(file);
        if (separator == EOF) {
            break;
        }
    }

    clock_t end_time = clock();

    fclose(file);

    double time_taken = ((double)(end_time - start_time)) / CLOCKS_PER_SEC;

    printf("Counter: %lld\n", global_counter);
    printf("Time: %f s\n", time_taken);

    return 0;
}