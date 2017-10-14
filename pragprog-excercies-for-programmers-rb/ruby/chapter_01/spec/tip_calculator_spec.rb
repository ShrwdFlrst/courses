require "./tip_calculator"

RSpec.describe TipCalculator do
  describe '#calculate_tip' do
    it 'returns the percentage of the total' do
      expect(TipCalculator.new.calculate_tip(11.25, 15)).to eq(1.69)
    end
  end

  describe '#calculate_total' do
    it 'returns the sum of the bill and the tip' do
      expect(TipCalculator.new.calculate_total(11.25, 1.69)).to eq(12.94)
    end
  end

  describe '#prompt_bill' do
    it 'ask the user for the bill' do
      expect{TipCalculator.new.prompt_bill}.to output("What is the bill?\n").to_stdout
    end
  end
end
