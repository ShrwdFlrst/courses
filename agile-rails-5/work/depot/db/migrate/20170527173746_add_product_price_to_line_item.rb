class AddProductPriceToLineItem < ActiveRecord::Migration[5.0]
  def self.up
    add_column :line_items, :price, :decimal,
      precision: 8, scale: 2, default: 0

    LineItem.all.each do |item|
      item.update_attribute :price, item.product.price
    end
  end

  def self.down
    remove_column :line_items, :price
  end
end
