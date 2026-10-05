<?php

use Illuminate\Foundation\Application;
use Illuminate\Foundation\Configuration\Exceptions;
use Illuminate\Foundation\Configuration\Middleware;
use Illuminate\Auth\AuthenticationException;
use Illuminate\Http\Request;
use Illuminate\Validation\ValidationException;
use App\Http\Middleware\PeranAdmin;
use Laravel\Sanctum\Exceptions\MissingAbilityException;
use Laravel\Sanctum\Http\Middleware\CheckAbilities;
use Symfony\Component\HttpKernel\Exception\NotFoundHttpException;

return Application::configure(basePath: dirname(__DIR__))
    ->withRouting(
        web: __DIR__.'/../routes/web.php',
        api: __DIR__.'/../routes/api.php',
        commands: __DIR__.'/../routes/console.php',
        health: '/up',
    )
    ->withMiddleware(function (Middleware $middleware): void {
        $middleware->alias([
            'ability' => CheckAbilities::class,
            'peran.admin' => PeranAdmin::class,
        ]);
    })
    ->withExceptions(function (Exceptions $exceptions): void {
        $exceptions->render(function (
            AuthenticationException $exception,
            Request $request
        ) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Token tidak valid atau belum dikirim.',
                ], 401);
            }
        });

        $exceptions->render(function (
            MissingAbilityException $exception,
            Request $request
        ) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Token tidak memiliki kemampuan yang cukup.',
                ], 403);
            }
        });

        $exceptions->render(function (
            NotFoundHttpException $exception,
            Request $request
        ) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Sumber daya tidak ditemukan',
                ], 404);
            }
        });

        $exceptions->render(function (
            ValidationException $exception,
            Request $request
        ) {
            if ($request->is('api/*')) {
                return response()->json([
                    'sukses' => false,
                    'pesan' => 'Data yang dikirim tidak valid',
                    'galat' => $exception->errors(),
                ], 422);
            }
        });
    })->create();
