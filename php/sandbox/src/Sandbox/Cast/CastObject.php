<?php declare(strict_types=1);

namespace Sandbox\Cast;

use InvalidArgumentException;

final class CastObject extends ParentObject
{
    public string $prop;

    public function __construct(string $prop)
    {
        $this->prop = $prop;
    }

    public static function cast($obj): self
    {
        if (!($obj instanceof self)) {
            throw new InvalidArgumentException("{$obj} is not instance of CastObject");
        }
        return $obj;
    }
}