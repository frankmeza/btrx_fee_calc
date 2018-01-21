Things to accomplish

# Tier 1
- come up with a sum of fees paid from trading bittrex in 2017
- problem is the moving value of BTC on a given day

## Tier 2
- decide how to record BTC price per day
    - in memory?
    - in DB?

### Tier 3
- readers in place
    - reading everything: each row as a slice

### next steps
- get just the low values from the hist data, with the date
    - capture within a struct