<?php

namespace App\Http\Requests;

use Illuminate\Foundation\Http\FormRequest;

class UpdateMatakuliahRequest extends FormRequest
{
    public function authorize(): bool
    {
        return true;
    }

    public function rules(): array
    {
        $matakuliah = $this->route('matakuliah');
        $id = is_object($matakuliah) ? $matakuliah->id : $matakuliah;

        return [
            'kode' => [
                'sometimes',
                'string',
                'max:20',
                'unique:matakuliahs,kode,' . $id,
            ],
            'nama' => ['sometimes', 'string', 'max:150'],
            'sks' => ['sometimes', 'integer', 'min:1', 'max:6'],
            'semester' => ['sometimes', 'integer', 'min:1', 'max:14'],
        ];
    }
}