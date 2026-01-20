- https://www.investopedia.com/calculate-principal-and-interest-5211981
- https://www.investopedia.com/terms/a/amortization.asp
- https://en.wikipedia.org/wiki/Amortization_calculator

- [Derivation of Loan/ Mortgage Monthly Payment Formula](https://youtu.be/iilFXMHKkZQ?si=igum12B4wlrCNunf)
- [Mortgage Monthly Payment Formula Derived](https://youtu.be/FUDW3kkzNYA?si=R5HrPlcPYUfem4Pd)
- [How to Create a Loan Amortization Schedule in Google Sheets/MS Excel](https://youtu.be/6g0OrQrVdJY?si=-a-ZSxmTa1bJAACr)
- [PMT()](https://support.google.com/docs/answer/3093185?hl=en)
- [IPMT()](https://support.google.com/docs/answer/3093175?hl=en)

Steps to Calculate Manually (or in a Spreadsheet) 

1. Gather Loan Details:
    - P (Principal): Total loan amount.
    - r (Annual Interest Rate): Convert to a monthly rate 
        
        (e.g., 6% becomes 0.06 / 12 = 0.005)
    - n (Total Number of Payments): Loan term in years * 12.
1. Calculate Your Fixed Monthly Payment (A):
    - Use the formula:

        $A=P\dfrac{r(1 + r)^n}{(1+r)^n -1}$
        
    - Or, use Excel's =PMT(rate, nper, pv) function (e.g., =PMT(0.06/12, 360, -300000)) to get your constant payment.
1. Build Your Schedule (Month by Month):
    - Month 1:Interest Paid: Loan Balance * Monthly Rate (e.g., $300,000 * 0.005 = $1,500).
    - Principal Paid: Monthly Payment - Interest Paid (e.g., $1,995.91 - $1,500 = $495.91).
    - New Balance: Loan Balance - Principal Paid (e.g., $300,000 - $495.91 = $299,504.09).
    - Month 2 & Beyond: Use the New Balance from the previous month as the Loan Balance for the current month and repeat the process. 
