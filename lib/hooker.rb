require 'sinatra'
require 'json'

class Hooker < Sinatra::Application
  post '/payload' do
    puts "I got some JSON: #{JSON.pretty_generate(JSON.parse(request.body.read))}"
  end
end
