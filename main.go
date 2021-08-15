package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (

	entered        int
	surname        string
	accountNum     int
	name           string
	number         float64
	Balance        [1000]float64
	FirstLast_name [1000][2]string
	Password       [1000]string
	Username       [1000]string
	user_name      string
	Transactions   [1000][30]float64
	accountID      int
	pass_word      string
	Fullname       string
	num            int
	picks          int
	choice         int
	s              string
	emailID        string
	phoneID        int
	email_ID       [1000]string
	phone_ID       [1000]int
	temp           float64

	array1   [7]string
	array2   [3]string
	phonenum string
	bal      float64
	i1       int

	TransactType [1000][30]string
	TransactiNum [1000][30]float64
	stemp        string
)

func main() {
	//setup transactiontype
	Setup()
start:
	fmt.Println("\t\t\t\t_____________________________________________________")
	fmt.Println("\t\t\t\t\t\tBANK MANAGEMENT SYSTEM")
	fmt.Println("MAIN MENU:")
	fmt.Println("1.Create Account")
	fmt.Println("2.Login")
	fmt.Println("3.Save")
	fmt.Println("4.Load")
	fmt.Println("5.Exit")
	fmt.Println("")
	fmt.Scanln(&picks)
	if picks < 0 || picks > 5 {
		fmt.Println("Please re-enter")
		goto start
	}
	fmt.Println("You Chose!", picks)
	switch picks {
	case 1:
		fmt.Println("")
		createAccount()
		goto start
	case 2:
		fmt.Println("")
		Login()
		goto start
	case 3:
		fmt.Println("")
		saveData()
		goto start
	case 4:
		fmt.Println("")
		loadData()
		goto start
	case 5:
		break

	}

}

func createAccount() {
	accountID++
	fmt.Println("Enter Firstname:")
	fmt.Scanln(&name)
	FirstLast_name[accountID][0] = name
	fmt.Println("Enter Lastname:")
	fmt.Scanln(&surname)
	FirstLast_name[accountID][1] = surname
	fmt.Println("Enter a Email:")
	fmt.Scanln(&emailID)
	email_ID[accountID] = emailID
	fmt.Println("Enter Phone Number:")
	fmt.Scanln(&phoneID)
	phone_ID[accountID] = phoneID
	fmt.Println("Enter a Username:")
	fmt.Scanln(&user_name)
	Username[accountID] = user_name
way1:
	fmt.Println("Enter Password:")
	fmt.Scanln(&pass_word)
	if len(pass_word) >= 8 && len(pass_word) < 17 {
		Password[accountID] = pass_word
	} else {
		fmt.Println("")
		fmt.Println("Must be at 8-16 characters")
		goto way1
	}
	fmt.Println("")
	fmt.Println("Your Username is:", user_name)
	fmt.Println("Your Name is ", name, surname)
	fmt.Println("Your email is: ", emailID)
	fmt.Println("Your Phone number is: ", phoneID)
	fmt.Println("Your Account number is ", accountID)
	fmt.Println("Congratulations! you created an account!")
}

func AccountAccess() {
	Fullname = FirstLast_name[num][0] + " " + FirstLast_name[num][1]
Again:
	fmt.Println("")
	fmt.Println("Would you like to: ")
	fmt.Println("1.Withdraw")
	fmt.Println("2.Deposit")
	fmt.Println("3.Account Details")
	fmt.Println("4.Transaction History")
	fmt.Println("5. Update Information")
	fmt.Println("6.Logout")
	fmt.Scanln(&choice)
	if choice < 0 || choice > 6 {
		fmt.Println("Please re-enter")
		goto Again
	}
	fmt.Println("\nYou Chose!", choice)
	switch choice {
	case 1:
		Withdraw()
		goto Again
	case 2:
		Deposit()
		goto Again
	case 3:
		AccountDetails()
		goto Again
	case 4:
		Transactionhistory()
		goto Again
	case 5:
		UpdateInfo()
		goto Again
	case 6:
		fmt.Println("Returning to main menu")
	}
}

func AccountDetails() {
	fmt.Println("Your Name is", FirstLast_name[num][0]+" "+FirstLast_name[num][1])
	fmt.Println("Your Account number is", num)
	fmt.Println("The Balance is", Balance[num])
	fmt.Println("Your email", email_ID[num])
	fmt.Println("Your Phone Number is", phone_ID[num])
}

func Transactionhistory() {
	fmt.Println("Your account number is: ", num)
	fmt.Println("Recent Transactions: ")
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("|Type    \t |Amount  \t\t |Balance \t\t|")
	fmt.Println("-----------------------------------------------------------------")
	for i := 0; i < 30; i++ {
		if TransactType[num][i] == "----" {
			//no history, then skip the rest
		} else {
			fmt.Print("|")
			fmt.Print(TransactType[num][i])
			fmt.Print("\t |")
			if TransactType[num][i] == "Deposit" {
				fmt.Print("+")
			} else {
				fmt.Print("-")
			}
			fmt.Print(TransactiNum[num][i])
			fmt.Print("\t\t\t |")
			fmt.Print(Transactions[num][i])
			fmt.Println("\t\t\t|")
		}

	}
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("")
}

func Deposit() {
	fmt.Println("Enter amount: ")
	fmt.Scanln(&number)
	Balance[num] = Balance[num] + number
	for i := 28; i >= 0; i-- {
		//Transaction
		temp = Transactions[num][i]
		Transactions[num][i+1] = temp
		//Type of Transaction
		stemp = TransactType[num][i]
		TransactType[num][i+1] = stemp
		//Transaction Amount
		temp = TransactiNum[num][i]
		TransactiNum[num][i+1] = temp
	}
	//Record Transaction history
	Transactions[num][0] = Balance[num]
	//Record Transaction Type
	TransactType[num][0] = "Deposit"
	//Record Transaction Amount
	TransactiNum[num][0] = number
}

func Withdraw() {
	fmt.Println("Enter amount: ")
	fmt.Scanln(&number)
	Balance[num] = Balance[num] - number
	for i := 28; i >= 0; i-- {
		//Transaction
		temp = Transactions[num][i]
		Transactions[num][i+1] = temp
		//Type of Transaction
		stemp = TransactType[num][i]
		TransactType[num][i+1] = stemp
		//Transaction Amount
		temp = TransactiNum[num][i]
		TransactiNum[num][i+1] = temp
	}
	//Record Transaction history
	Transactions[num][0] = Balance[num]
	//Record Transaction Type
	TransactType[num][0] = "Withdraw"
	//Record Transaction Amount
	TransactiNum[num][0] = number
}

func Login() {
allover:
	name = s
	surname = s
	pass_word = s
	fmt.Println("To return to main menu enter 0 for one/any/all of the following")
	fmt.Println("Enter Account number:")
	fmt.Scanln(&accountNum)
	num = accountNum
	if num >= 0 && num < 1001 {
		goto allover1
	} else {
		goto allover
	}
allover1:
	fmt.Println("Enter Username:")
	fmt.Scanln(&name)
	if name != s {
		goto allover2
	} else {
		goto allover1
	}
allover2:
	fmt.Println("Enter Password:")
	fmt.Scanln(&pass_word)
	if name == Username[num] && pass_word == Password[num] {
		AccountAccess()
	} else if accountNum == 0 || name == "0" || pass_word == "0" {
		fmt.Println("")
		fmt.Println("Returning to Main Menu")
	} else {
		fmt.Println("Account not found, please enter correct credentials")
		fmt.Println("")
		goto allover
	}
}

func UpdateInfo() {

allover:
	name = s
	surname = s
	pass_word = s
	fmt.Println("To return to main menu enter 0 for one/any/all of the following")
	fmt.Println("Enter Account number:")
	fmt.Scanln(&accountNum)
	num = accountNum
	if num >= 0 && num < 1001 {
		goto allover1
	} else {
		goto allover
	}
allover1:
	fmt.Println("Enter Username:")
	fmt.Scanln(&name)
	if name != s {
		goto allover2
	} else {
		goto allover1
	}
allover2:
	fmt.Println("Enter Password:")
	fmt.Scanln(&pass_word)
	if name == Username[num] && pass_word == Password[num] {

	} else if accountNum == 0 || name == "0" || pass_word == "0" {
		fmt.Println("")
		fmt.Println("Returning to Main Menu")
	} else {
		fmt.Println("Account not found, please enter correct credentials")
		fmt.Println("")
		goto allover
	}
	fmt.Println("Enter New Firstname:")
	fmt.Scanln(&name)
	FirstLast_name[accountID][0] = name
	fmt.Println("Enter New Lastname:")
	fmt.Scanln(&surname)
	FirstLast_name[accountID][1] = surname
	fmt.Println("Enter new a Email:")
	fmt.Scanln(&emailID)
	email_ID[accountID] = emailID
	fmt.Println("Enter new Phone Number:")
	fmt.Scanln(&phoneID)
	phone_ID[accountID] = phoneID
	fmt.Println("Enter new Username:")
	fmt.Scanln(&user_name)
	Username[accountID] = user_name
way1:
	fmt.Println("Enter new Password:")
	fmt.Scanln(&pass_word)
	if len(pass_word) >= 8 && len(pass_word) < 17 {
		Password[accountID] = pass_word
	} else {
		fmt.Println("")
		fmt.Println("Must be at 8-16 characters")
		goto way1
	}
	fmt.Println("")
	fmt.Println("Your Username is:", user_name)
	fmt.Println("Your Name is ", name, surname)
	fmt.Println("Your email is: ", emailID)
	fmt.Println("Your Phone number is: ", phoneID)
	fmt.Println("Your Account number is ", accountID)
	fmt.Println("Congratulations! you have changed an account!")
}

/***************************************************************************************************************************/

func saveData() {
	//Create/Overwrite AccountData text document
	f, err := os.Create("AccountData.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Close file after function is finished
	defer f.Close()

	//Create a string type slice
	var slice1 []string = array1[:]

	//Write and Store the information for each Account (1000)
	for i := 1; i < 1000; i++ {
		//Element 0 Firstname
		slice1[0] = FirstLast_name[i][0]
		//Element 1 Lastname
		slice1[1] = FirstLast_name[i][1]

		//Element 2 Phone Number
		//convert from type int to type string
		phonenum := strconv.Itoa(phone_ID[i])
		slice1[2] = phonenum

		//Element 3 Email Address
		slice1[3] = email_ID[i]

		//Element 4 Username
		slice1[4] = Username[i]

		//Element 5 Password
		slice1[5] = Password[i]

		//Element 6 Balance
		//convert from type float64 to type string
		bal := strconv.FormatFloat(Balance[i], 'f', -1, 64)
		slice1[6] = bal

		//Write the contents of the slice in one line
		for _, slice1 := range slice1 {
			_, err := f.WriteString(slice1 + " ")
			if err != nil {
				log.Fatal(err)
			}

		}

		//Write a new line
		_, err := f.WriteString("\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Saving...")
	saveTData()
}

func loadData() {

	//Open AccountData text document
	file, err := os.Open("AccountData.txt")
	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	//Close the file
	file.Close()

	//Print each of the slice values
	for _, each_ln := range text {
		//Write the contents of the slice in one line
		words := strings.Fields(each_ln) //each word in a line
		if len(words) > 2 {
			i1++
			FirstLast_name[i1][0] = words[0]
			FirstLast_name[i1][1] = words[1]

			//phone number
			str1 := words[2]
			i2, err := strconv.Atoi(str1)
			if err == nil {
				phone_ID[i1] = i2
			}

			email_ID[i1] = words[3]
			Username[i1] = words[4]
			Password[i1] = words[5]

			//balance
			str2 := words[6]
			f64, err := strconv.ParseFloat(str2, 64)
			if err == nil {
				Balance[i1] = f64
			}
			//fmt.Println(words, len(words)) // [one two three four...]
			accountID = i1
		} else {
			//fmt.Println(words, len(words)) // [one two three four...]
		}
	}

	fmt.Println("Loading...")
	loadTData()
}

/***************************************************************************************************************************/

func saveTData() {
	//Create/Overwrite TransactionData text document
	f, err := os.Create("TransactionData.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Close file after function is finished
	defer f.Close()

	//Create a string type slice
	var slice2 []string = array2[:]

	//Write and Store the transaction info (30) for each Account (1000)
	for i := 1; i < 1000; i++ {
		for j := 0; j < 30; j++ {

			//Element 0 Transaction history
			//convert from type float64 to type string
			bal := strconv.FormatFloat(Transactions[i][j], 'f', -1, 64)
			slice2[0] = bal

			//Element 1 Type of Transaction
			slice2[1] = TransactType[i][j]

			//Element 2 Transaction Amount
			//convert from type float64 to type string
			bal = strconv.FormatFloat(TransactiNum[i][j], 'f', -1, 64)
			slice2[2] = bal

			//Write the contents of the slice in one line
			for _, slice2 := range slice2 {
				_, err := f.WriteString(slice2 + " ")
				if err != nil {
					log.Fatal(err)
				}

			}

			//Write a new line
			_, err := f.WriteString("\n")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("Data has been saved!")
}

func loadTData() {
	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open("TransactionData.txt")

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.
	i := 1
	j := 0

	for _, each_ln := range text {
		if j == 29 {
			j = 0
			i++
		} else if i == 999 {
			break
		}
		words := strings.Fields(each_ln) //each word in a line
		if len(words) > 2 {

			//Transactions
			str2 := words[0]
			f64, err := strconv.ParseFloat(str2, 64)
			if err == nil {
				Transactions[i][j] = f64
			}

			//Type
			TransactType[i][j] = words[1]

			//Amount
			str2 = words[2]
			f64, err = strconv.ParseFloat(str2, 64)
			if err == nil {
				TransactiNum[i][j] = f64
			}

			//fmt.Println(words, len(words)) // [one two three]

		} else {
			//fmt.Println(words, len(words)) // [one two three]
		}
		j++
	}
	fmt.Println("Loading complete!")

}

func Setup() {
	for i := 0; i < 1000; i++ {
		for j := 0; j < 30; j++ {
			TransactType[i][j] = "----"
		}
	}
}
