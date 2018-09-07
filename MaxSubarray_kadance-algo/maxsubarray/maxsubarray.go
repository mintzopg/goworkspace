package maxsubarray

//  Tο πρόβλημα μέγιστου αθροίσματος υποακολουθίας (Αγγλικά: Maximum subarray problem) είναι ένα πρόβλημα όπου έχουμε μια σειρά (πίνακα) από θετικούς και αρνητικούς αριθμούς μια ψάχνουμε την συνεχή μεγαλύτερη υποσειρά (υποπίνακα) αριθμών όπου έχουμε το μεγαλύτερο άθροισμα των στοιχείων. Για παράδειγμα στην σειρά των αριθμών −2, 1, −3, 4, −1, 2, 1, −5, 4 η υποσειρά αριθμών με μέγιστο άθροισμα είναι η 4, −1, 2, 1, με άθροισμα 6.

// https://el.wikipedia.org/wiki/Πρόβλημα_μέγιστου_αθροίσματος_υποακολουθίας

// Sequence function (arra []int) ([]int, int)
func Sequence(arr []int) ([]int, int) {
	// αρχικοποίηση
	maxsofar := arr[0]
	maxendinghere := arr[0]

	// Αυτές οι μεταβλητές μπορούν να χρησιμοποιηθούν για να έχουμε τους δείκτες αρχής και τέλους της υποσειράς
	var begin, begintemp, end int

	// Βρες την υποσειρά κάνοντας επανάληψη μέχρι το τέλος του πίνακα
	for i := 1; i < len(arr); i++ {
		// υπολογισμός maxendinghere
		if maxendinghere < 0 {
			maxendinghere = arr[i]
			begintemp = i
		} else {
			maxendinghere += arr[i]
		}
		// υπολογισμός maxsofar
		if maxendinghere >= maxsofar {
			maxsofar = maxendinghere
			begin = begintemp
			end = i
		}
	}

	subarr := arr[begin : end+1]
	return subarr, maxsofar
}
