# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: plane/plane.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("plane/plane.proto", :syntax => :proto3) do
    add_message "Coordinates" do
      optional :x, :int64, 1
      optional :y, :int64, 2
    end
    add_message "Point" do
      optional :coordinates, :message, 1, "Coordinates"
      optional :count, :uint64, 2
    end
    add_message "Bounds" do
      optional :min, :message, 1, "Coordinates"
      optional :max, :message, 2, "Coordinates"
    end
  end
end

Coordinates = Google::Protobuf::DescriptorPool.generated_pool.lookup("Coordinates").msgclass
Point = Google::Protobuf::DescriptorPool.generated_pool.lookup("Point").msgclass
Bounds = Google::Protobuf::DescriptorPool.generated_pool.lookup("Bounds").msgclass
