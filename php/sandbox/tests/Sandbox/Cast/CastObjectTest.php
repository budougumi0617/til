<?php declare(strict_types=1);

namespace Tests\Sandbox\Cast;

use PHPUnit\Framework\TestCase;
use Sandbox\Cast\CastObject;
use Sandbox\Cast\ParentObject;

final class CastObjectTest extends TestCase
{

    private static function getProp(ParentObject $obj): string
    {
        $convert = fn($orig): CastObject => $orig;
        return $convert($obj)->prop;
    }

    private static function getObject(string $prop): ParentObject
    {
        return new CastObject($prop);
    }

    public function testCast()
    {
        $data = 'Data';
        $obj = self::getObject($data);

        $cast = fn($orig): CastObject => $orig;
        $casted = $cast($obj);

        if ($obj instanceof CastObject) {
            $casted = CastObject::cast($obj);
        }

        $this->assertSame($casted->prop, $data);
    }
}
