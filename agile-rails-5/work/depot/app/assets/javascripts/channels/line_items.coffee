App.line_items = App.cable.subscriptions.create "LineItemsChannel",
  connected: ->
    # Called when the subscription is ready for use on the server

  disconnected: ->
    # Called when the subscription has been terminated by the server

  received: (data) ->
    $('#cart').html(data.html)
    $('#current_item').css({'background-color': '#88ff88'}).
      animate({'background-color': '#114411'}, 1000)
