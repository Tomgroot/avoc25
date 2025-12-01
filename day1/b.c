#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {

    if (argc != 2) {
        fprintf(stderr, "Usage: %s <input_file>\n", argv[0]);
        return 1;
    }

    FILE *file = fopen(argv[1], "r");
    if (file == NULL) {
        perror("Error opening file");
        return 1;
    }

    int current_value = 50;
    int zero_count = 0;

    char direction;
    int amount;
    char line_buffer[100];

    while (fgets(line_buffer, sizeof(line_buffer), file)) {
        if (sscanf(line_buffer, "%c%d", &direction, &amount) != 2) {
            fprintf(stderr, "Error parsing line: %s\n", line_buffer);
            return 1;
        }

        for (int i = 0; i < amount; i++) {
            if (direction == 'L') {
                current_value--;
                if (current_value < 0) {
                    current_value = 99;
                }
            } else if (direction == 'R') {
                current_value++;
                if (current_value > 99) {
                    current_value = 0;
                }
            }

            if (current_value == 0) {
                zero_count++;
            }
        }

    }

    fclose(file);
    
    printf("Final Value: %d\n", current_value);
    printf("Total times value hit 0: %d\n", zero_count);

    return 0;
}