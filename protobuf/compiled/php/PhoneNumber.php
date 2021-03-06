<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: address_book/address_book.proto

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;
use Google\Protobuf\Internal\GPBWrapperUtils;

/**
 * Generated from protobuf message <code>PhoneNumber</code>
 */
class PhoneNumber extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string number = 1;</code>
     */
    private $number = '';
    /**
     * Generated from protobuf field <code>.PhoneType type = 2;</code>
     */
    private $type = 0;
    /**
     * Generated from protobuf field <code>string extension = 3;</code>
     */
    private $extension = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $number
     *     @type int $type
     *     @type string $extension
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\AddressBook\AddressBook::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string number = 1;</code>
     * @return string
     */
    public function getNumber()
    {
        return $this->number;
    }

    /**
     * Generated from protobuf field <code>string number = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.PhoneType type = 2;</code>
     * @return int
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * Generated from protobuf field <code>.PhoneType type = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setType($var)
    {
        GPBUtil::checkEnum($var, \PhoneType::class);
        $this->type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string extension = 3;</code>
     * @return string
     */
    public function getExtension()
    {
        return $this->extension;
    }

    /**
     * Generated from protobuf field <code>string extension = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setExtension($var)
    {
        GPBUtil::checkString($var, True);
        $this->extension = $var;

        return $this;
    }

}

