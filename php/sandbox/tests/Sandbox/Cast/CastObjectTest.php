<?php declare(strict_types=1);

namespace Tests\Sandbox\Cast;

use PHPUnit\Framework\TestCase;
use Sandbox\Cast\CastObject;
use Sandbox\Cast\ParentObject;

final class CastObjectTest extends TestCase
{

    private function getProp(ParentObject $obj): string
    {
        $convert = fn($orig): CastObject => $orig;
        return $convert($obj)->prop;
    }

    public function testCastBasicLit()
    {
        $obj = new CastObject();
        $casted = CastObject::castBasicLit($obj);
        $this->assertSame($casted->prop, CastObject::class);
    }
}
