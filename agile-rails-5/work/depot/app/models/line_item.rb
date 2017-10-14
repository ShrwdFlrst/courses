class LineItem < ApplicationRecord
  belongs_to :product
  belongs_to :cart

  def total_price
    price * quantity
  end

  def decrement
    if quantity == 1
      self.destroy
    else
      self.quantity = quantity - 1
      self.save
    end
  end
end
