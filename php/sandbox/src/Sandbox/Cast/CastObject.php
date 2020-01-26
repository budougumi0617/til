<?php declare(strict_types=1);

namespace Sandbox\Cast;

use InvalidArgumentException;

final class CastObject
{
    public static function castBasicLit($obj): self
    {
        if (!($obj instanceof self)) {
            throw new InvalidArgumentException("{$obj} is not instance of CastObject");
        }
        return $obj;
    }
}