Rails.application.routes.draw do
  resources :line_items
  put 'line_items/:id/decrement', to: 'line_items#decrement', as: 'decrement_line_item'
  resources :carts
  root 'store#index', as: 'store_index'

  resources :products

  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
end
