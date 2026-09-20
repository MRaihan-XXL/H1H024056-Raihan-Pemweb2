<!DOCTYPE html>
<html lang="id">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>@yield('judul', 'Pemweb II')</title>
    <style>
        body {
            margin: 0;
            font-family: Arial, sans-serif;
            background: #f5f5f3;
            color: #1a1a1a;
        }

        .wrap {
            max-width: 980px;
            margin: 0 auto;
            padding: 24px 16px 40px;
        }

        .topbar {
            border-bottom: 1px solid #e3e0db;
            background: #fff;
        }

        .topbar .wrap {
            padding-top: 16px;
            padding-bottom: 16px;
        }

        .brand {
            text-decoration: none;
            color: #111;
            font-size: 12px;
            letter-spacing: 0.14em;
            text-transform: uppercase;
            font-weight: 700;
        }

        .card {
            background: #fff;
            border: 1px solid #e7e2dc;
            border-radius: 14px;
            padding: 20px;
        }

        h1, h2, h3, p {
            margin-top: 0;
        }

        h1 {
            font-size: 28px;
            margin-bottom: 10px;
        }

        .muted {
            color: #666;
            font-size: 12px;
            text-transform: uppercase;
            letter-spacing: 0.12em;
            margin-bottom: 8px;
        }

        table {
            width: 100%;
            border-collapse: collapse;
            font-size: 14px;
        }

        th, td {
            padding: 12px 10px;
            border-bottom: 1px solid #eee;
            text-align: left;
        }

        th {
            color: #666;
            font-size: 12px;
            text-transform: uppercase;
            letter-spacing: 0.08em;
        }

        .btn {
            display: inline-block;
            padding: 8px 14px;
            border: 1px solid #111;
            border-radius: 999px;
            text-decoration: none;
            color: #111;
            background: #fff;
            font-size: 13px;
        }

        .btn.primary {
            background: #111;
            color: #fff;
        }

        .toolbar {
            display: flex;
            gap: 10px;
            align-items: end;
            flex-wrap: wrap;
            margin-bottom: 20px;
        }

        .field {
            flex: 1 1 260px;
        }

        .field input {
            box-sizing: border-box;
            width: 100%;
            padding: 9px 11px;
            border: 1px solid #d8d3cd;
            border-radius: 8px;
            font: inherit;
        }

        .badge {
            display: inline-block;
            padding: 4px 8px;
            border-radius: 6px;
            font-size: 12px;
            font-weight: 700;
        }

        .badge.warning {
            background: #f4df9b;
            color: #493b0b;
        }

        .badge.success {
            background: #cfe8d4;
            color: #1d5729;
        }

        .table-wrap {
            overflow-x: auto;
        }

        .row {
            display: grid;
            grid-template-columns: 1fr 1.5fr;
            gap: 20px;
        }

        .box {
            background: #fff;
            border: 1px solid #e7e2dc;
            border-radius: 14px;
            padding: 18px;
        }

        .label {
            color: #666;
            font-size: 12px;
            display: block;
            margin-bottom: 4px;
        }

        .value {
            font-weight: 700;
        }

        @media (max-width: 700px) {
            .row {
                grid-template-columns: 1fr;
            }

            th, td {
                white-space: nowrap;
            }
        }
    </style>
</head>
<body>
    <nav class="topbar">
        <div class="wrap">
            <a class="brand" href="{{ url('/') }}">Pemweb II</a>
        </div>
    </nav>

    <main class="wrap">
        @yield('konten')
    </main>
</body>
</html>