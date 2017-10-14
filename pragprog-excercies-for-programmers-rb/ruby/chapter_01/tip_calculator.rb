class TipCalculator
    def calculate_tip (bill, percentage)
        ((bill / 100) * percentage).round(2)
    end

    def calculate_total (bill, tip)
        bill + tip
    end

    def prompt_bill
        puts "What is the bill?"
    end
end
