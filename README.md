Things to accomplish

# Tier 1
- come up with a sum of fees paid from trading bittrex in 2017
- problem is the moving value of BTC on a given day

## Tier 2
- readers in place
    - reading everything: each row as a slice
- structs in place
    - capturing fee{date, commissionPaid}
    - capturing price{date, price}

### next steps
- convert commissionPaid in BTC into USD on that day
    - make uniform the date structures between fees and price
    - record `feePaid = commissionPaid / price`