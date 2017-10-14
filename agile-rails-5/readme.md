    cd /c/Users/fargo/dev/agile-rails-5; vagrant up
    cd /c/Users/fargo/dev/agile-rails-5; vagrant ssh

    bin/rails s -b 0.0.0.0

    bin/rails g channel line_items

    bin/rails g controller Say experiments
    bin/rails generate scaffold Cart
    bin/rails generate scaffold LineItem product:references
    bin/rails generate scaffold LineItem product:references cart:belongs_to --force
    bin/rails generate scaffold Product title:string description:text image_url:string price:decimal --force
    bin/rails dev:cache

    bin/rails test
    bin/rails test:controllers
    bin/rails test:models

    bin/rails db:fixtures:load
    bin/rails db:seed RAILS_ENV=test
    bin/rails db:seed
    bin/rails db:drop

    bin/rails c RAILS_ENV=test

    rm -rf db/development.sqlite3 db/test.sqlite3
    bin/rails db:migrate:redo RAILS_ENV=development
    bin/rails db:migrate RAILS_ENV=development
    bin/rails db:migrate RAILS_ENV=test
    bin/rails db:migrate:status
    bin/rails db:rollback
    bin/rails generate migration copy_product_price_into_line_item
    bin/rails generate migration combine_items_in_cart

    bin/rails log:clear LOGS=test
