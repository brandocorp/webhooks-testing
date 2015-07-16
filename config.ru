require 'rubygems'
require 'bundler/setup'

Bundler.require

require 'sinatra'
require_relative 'lib/hooker'

# Set the environment to :production on production
set :environment, ENV['RACK_ENV'].to_sym

# And fire the application.
run Hooker
