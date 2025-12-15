#include <bits/stdc++.h>
using namespace std;

int main() {
    ifstream file("input.txt");

    if (!file.is_open()) {
        cerr << "Error: Could not open input.txt\n";
        return 1;
    }

    std::string line;

    int ans = 50;

    int count = 0;

    while (getline(file, line)) {
        if (line.empty())
            continue;

        
        char letter = line[0];

        
        int number = stoi(line.substr(1));

        number = number % 100;

        if(letter == 'R'){
            ans = (ans+number)%100;
        }
        else{
            ans = (ans - number + 100) % 100;
        }

        if(ans == 0){
            count++;
        }


    }

    cout << count << endl;

    

    file.close();
    return 0;
}
