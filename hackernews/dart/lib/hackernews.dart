@JsName('hackernews')
library hackernews;

import 'dart:async';

import 'package:js_wrapping/js_wrapping.dart';

part 'hackernews.g.dart';

final fromDynamic = DynamicCodec();
final fromMap = JsObjectAsMapCodec<dynamic>(fromDynamic);
final fromListMap = JsListCodec<Map>(fromMap);

abstract class _HackerNews implements JsInterface {
  factory _HackerNews(String baseUrl) => null;

  JsObject GetFeed(String name, num page);
  JsObject GetItem(String id);

  Future<List<Map>> getFeed(String name, num page) {
    final Completer c = Completer();
    final JsObject o = GetFeed(name, page);
    o.callMethod('then', [c.complete]);
    return c.future.then((v) {
      return fromListMap.decode(v);
    });
  }

  Future<Map> getItem(String id) {
    final Completer c = Completer();
    final JsObject o = GetItem(id);
    o.callMethod('then', [c.complete]);
    return c.future.then((v) {
      return fromMap.decode(v);
    });
  }
}
