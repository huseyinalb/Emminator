{Emminator} = require('./emminator')
class Hede extends Emminator
  func1: ->
    @emit "hi1"
  func2: ->
    @emit "hi2"
hede = new Hede

hede.on "hi1", -> console.log "here1"
hede.once "hi2", -> console.log "here2"
hede.ready -> console.log "ready"
hi2 = ->console.log "hi2"
hede.on "hi2", -> hi2
hede.func1()
hede.func1()
hede.func2()
hede.func2()
hede.off 'hi2', hi2
hede.emit 'hi2'
hede.emit 'ready'
hede.ready -> console.log "ready"
