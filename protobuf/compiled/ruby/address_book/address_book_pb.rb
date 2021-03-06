# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: address_book/address_book.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("address_book/address_book.proto", :syntax => :proto3) do
    add_message "PhoneNumber" do
      optional :number, :string, 1
      optional :type, :enum, 2, "PhoneType"
      optional :extension, :string, 3
    end
    add_message "Person" do
      optional :id, :int64, 1
      optional :name, :string, 2
      optional :email, :string, 3
      repeated :phones, :message, 4, "PhoneNumber"
      map :tags, :string, :string, 5
      optional :isBlocked, :bool, 6
      optional :age, :uint32, 8
    end
    add_message "AddressBook" do
      repeated :people, :message, 1, "Person"
    end
    add_enum "PhoneType" do
      value :MOBILE, 0
      value :HOME, 1
      value :WORK, 2
    end
  end
end

PhoneNumber = Google::Protobuf::DescriptorPool.generated_pool.lookup("PhoneNumber").msgclass
Person = Google::Protobuf::DescriptorPool.generated_pool.lookup("Person").msgclass
AddressBook = Google::Protobuf::DescriptorPool.generated_pool.lookup("AddressBook").msgclass
PhoneType = Google::Protobuf::DescriptorPool.generated_pool.lookup("PhoneType").enummodule
