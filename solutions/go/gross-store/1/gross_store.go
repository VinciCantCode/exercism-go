package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit := map[string]int{}
	unit["quarter_of_a_dozen"] = 3
	unit["half_of_a_dozen"] = 6
	unit["dozen"] = 12
	unit["small_gross"] = 120
	unit["gross"] = 144
	unit["great_gross"] = 1728
	return unit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, okUnit := units[unit]
	if okUnit == false {
		return false
	} else {
		bill[item] += units[unit]
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, okUnit := units[unit]
	_, okBill := bill[item]
	if okUnit == false {
		return false
	} else {
		if okBill == false {
			return false
		} else {
			if bill[item]-units[unit] < 0 {
				return false
			} else if bill[item]-units[unit] == 0 {
				delete(bill, item)
			} else {
				bill[item] -= units[unit]
			}
			return true
		}
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, okBill := bill[item]
	if okBill == false {
		return 0, false
	} else {
		return bill[item], true
	}
}
