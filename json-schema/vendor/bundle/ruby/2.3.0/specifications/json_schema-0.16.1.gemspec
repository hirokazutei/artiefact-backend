# -*- encoding: utf-8 -*-
# stub: json_schema 0.16.1 ruby lib

Gem::Specification.new do |s|
  s.name = "json_schema".freeze
  s.version = "0.16.1"

  s.required_rubygems_version = Gem::Requirement.new(">= 0".freeze) if s.respond_to? :required_rubygems_version=
  s.require_paths = ["lib".freeze]
  s.authors = ["Brandur".freeze]
  s.date = "2016-12-16"
  s.email = ["brandur@mutelight.org".freeze]
  s.executables = ["validate-schema".freeze]
  s.files = ["bin/validate-schema".freeze]
  s.homepage = "https://github.com/brandur/json_schema".freeze
  s.licenses = ["MIT".freeze]
  s.rubygems_version = "2.5.2.3".freeze
  s.summary = "A JSON Schema V4 and Hyperschema V4 parser and validator.".freeze

  s.installed_by_version = "2.5.2.3" if s.respond_to? :installed_by_version

  if s.respond_to? :specification_version then
    s.specification_version = 4

    if Gem::Version.new(Gem::VERSION) >= Gem::Version.new('1.2.0') then
      s.add_development_dependency(%q<ecma-re-validator>.freeze, ["~> 0.1"])
      s.add_development_dependency(%q<minitest>.freeze, ["~> 5.3"])
      s.add_development_dependency(%q<rake>.freeze, ["~> 10.3"])
      s.add_development_dependency(%q<byebug>.freeze, [">= 0"])
      s.add_development_dependency(%q<pry>.freeze, [">= 0"])
      s.add_development_dependency(%q<pry-byebug>.freeze, [">= 0"])
      s.add_development_dependency(%q<simplecov>.freeze, [">= 0"])
    else
      s.add_dependency(%q<ecma-re-validator>.freeze, ["~> 0.1"])
      s.add_dependency(%q<minitest>.freeze, ["~> 5.3"])
      s.add_dependency(%q<rake>.freeze, ["~> 10.3"])
      s.add_dependency(%q<byebug>.freeze, [">= 0"])
      s.add_dependency(%q<pry>.freeze, [">= 0"])
      s.add_dependency(%q<pry-byebug>.freeze, [">= 0"])
      s.add_dependency(%q<simplecov>.freeze, [">= 0"])
    end
  else
    s.add_dependency(%q<ecma-re-validator>.freeze, ["~> 0.1"])
    s.add_dependency(%q<minitest>.freeze, ["~> 5.3"])
    s.add_dependency(%q<rake>.freeze, ["~> 10.3"])
    s.add_dependency(%q<byebug>.freeze, [">= 0"])
    s.add_dependency(%q<pry>.freeze, [">= 0"])
    s.add_dependency(%q<pry-byebug>.freeze, [">= 0"])
    s.add_dependency(%q<simplecov>.freeze, [">= 0"])
  end
end
