<?php

namespace Tests\Sandbox\Cast;

use PHPUnit\Framework\TestCase;
use Sandbox\Cast\CastObject;

class CastObjectTest extends TestCase
{

    public function testCastBasicLit()
    {
        $obj = new CastObject();
        $casted = CastObject::castBasicLit($obj);
        $this->assertSame($obj, $casted);
    }
}
