require 'test_helper'

class CartTest < ActiveSupport::TestCase
  def setup
    @cart = Cart.create
    @product_ruby = products(:ruby)
    @product_one = products(:one)
  end

  test "add same products" do
    @cart.add_product(@product_ruby).save!
    @cart.add_product(@product_ruby).save!

    assert_equal 1, @cart.line_items.size
    assert_equal 2, @cart.line_items[0].quantity
    assert_equal @product_ruby.price * 2, @cart.total_price
  end

  test "add unique products" do
    @cart.add_product(@product_one).save!
    @cart.add_product(@product_ruby).save!

    assert_equal 2, @cart.line_items.size
    assert_equal 1, @cart.line_items[0].quantity
    assert_equal 1, @cart.line_items[1].quantity
    assert_equal @product_ruby.price + @product_one.price, @cart.total_price
  end
end
