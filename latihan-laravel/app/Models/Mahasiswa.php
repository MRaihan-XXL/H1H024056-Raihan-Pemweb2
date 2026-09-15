<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;
use Illuminate\Database\Eloquent\Relations\BelongsToMany;

class Mahasiswa extends Model
{
    protected $table = 'mahasiswas';

    protected $fillable = ['nim', 'nama', 'angkatan', 'program_studi'];

    public $timestamps = false;

    public function matakuliahs(): BelongsToMany
    {
        return $this->belongsToMany(Matakuliah::class, 'mahasiswa_matakuliah')
            ->withPivot('nilai')
            ->withTimestamps();
    }
}
