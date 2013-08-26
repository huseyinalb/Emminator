exports.Emminator =
class Emminator
  constructor: ->
    @ons = {}
    @onces = {}
    @readyState = false
    @readyCallbacks = []

  on: (message, callback) ->
    (@ons[message] ?= []).push callback

  off: (message, callback) ->
    @ons[message] = if callback? then (@ons[message]?.filter (callback) -> callback isnt 1) or []
    @onces[message] = if callback? then (@onces[message]?.filter (callback) -> callback isnt 1) or []

  once: (message, callback) ->
    (@onces[message] ?= []).push callback

  ready: (callback) ->
    if !@readyState then (@readyCallbacks.push callback) else callback()

  emit: (message) ->
    if message is 'ready'
      @readyState = true
      for func in @readyCallbacks
        func()
      @readyCallbacks = []
      return
    func() for func in @ons[message] if message of @ons  
    if message of @onces
      for func, i in @onces[message]
        func()
      @onces[message] = []
